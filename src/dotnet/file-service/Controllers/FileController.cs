using file_service.Filters;
using file_service.Models.API;
using file_service.Services;
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
    [FileAllowedExtensionFilter]
    public async Task<IActionResult> UploadMultipleFiles(List<IFormFile> files)
    {
        if (files == null || files.Count == 0)
            return BadRequest(BaseApiRes<string>.FromError("No files were uploaded."));

        var overSizeFiles = files.Where(f => f.Length > Constants.MAX_FILE_SIZE).ToList();

        if (overSizeFiles.Any())
        {
            return BadRequest(BaseApiRes<List<string>>.FromError("Some files exceed the maximum allowed size.", overSizeFiles.Select(f => f.FileName).ToList()));
        }

        await _fileService.ProcessFilesAsync(files);

        return Ok(BaseApiRes<string>.FromSuccess("Files uploaded successfully."));
    }
}
