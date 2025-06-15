using file_service.Filters;
using file_service.Models.API;
using file_service.Services;
using file_service.Extensions;
using Microsoft.AspNetCore.Mvc;
using file_service.Utils;
using file_service.Models;
using System.Threading.Tasks;
using System.Diagnostics;

namespace file_service.Controllers;

[ApiController]
[Route("[controller]")]
public class FileController : ControllerBase
{
    private readonly IFileService _fileService;

    public FileController(IFileService fileService)
    {
        _fileService = fileService;
    }

    [HttpPost]
    [FileTypeAllowedFilter]
    public async Task<IActionResult> UploadMultipleFiles(List<IFormFile> files)
    {
        if (files == null || files.Count == 0)
        {
            return BadRequest(BaseApiRes<string>.FromError("No files were uploaded."));
        }

        var overSizeFiles = files.Where(f => f.IsFileOverSize());

        if (overSizeFiles.Any())
        {
            return BadRequest(BaseApiRes<IEnumerable<string>>.FromError(
                "Some files exceed the maximum allowed size.",
                overSizeFiles.Select(f => f.FileName)
            ));
        }

        var processedFiles = await _fileService.ProcessFilesAsync(files);
        return Ok(BaseApiRes<IEnumerable<string>>.FromSuccess(processedFiles.Select(f => f.FileName), "Files uploaded successfully."));
    }

    [HttpGet("download/{fileId}")]
    public async Task<IActionResult> DownloadFile(Guid fileId)
    {
        var attachment = await _fileService.GetFileAsync(fileId);
        if (attachment == null)
        {
            return NotFound(BaseApiRes<string>.FromError("File not found."));
        }

        if (attachment.IsLocalFile)
        {

            var filePath = FileStorage.GetLocalFilePath(attachment.FileUrl);
            if (!System.IO.File.Exists(filePath))
            {
                return NotFound(BaseApiRes<string>.FromError("File does not exist on the server."));
            }
            var fileBytes = await System.IO.File.ReadAllBytesAsync(filePath);
            return File(fileBytes, attachment.ContentType, attachment.FileName);
        }
        return Ok(BaseApiRes<string>.FromSuccess(attachment.FileUrl, "File URL retrieved successfully."));
    }

    [HttpGet("{fileId}")]
    public async Task<IActionResult> GetFile(Guid fileId, [FromQuery] bool is_stream = false)
    {
        var attachment = await _fileService.GetFileAsync(fileId);

        if (attachment == null)
        {
            return NotFound(BaseApiRes<string>.FromError("File not found."));
        }

        if (is_stream)
        {
            if (!attachment.IsLocalFile)
            {
                return BadRequest(BaseApiRes<string>.FromError("Streaming is only supported for local files."));
            }

            var stream = new FileStream(
                FileStorage.GetLocalFilePath(attachment.FileUrl),
                FileMode.Open,
                FileAccess.Read,
                FileShare.Read
            );

            return File(stream, attachment.ContentType, attachment.FileName);
        }

        return Ok(BaseApiRes<Attachment>.FromSuccess(attachment, "File retrieved successfully."));
    }
}
