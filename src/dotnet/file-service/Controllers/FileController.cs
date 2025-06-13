using file_service.Filters;
using file_service.Models.API;
using file_service.Services;
using file_service.Extensions;
using Microsoft.AspNetCore.Mvc;

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

    [HttpGet("stream/{fileId}")]
    public IActionResult StreamFile(Guid fileId)
    {
        return Ok();
    }
}
