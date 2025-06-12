using file_service.Models.API;
using Microsoft.AspNetCore.Mvc.Filters;

namespace file_service.Filters;

public class FileAllowedExtensionFilter : Attribute, IAsyncActionFilter
{
    private readonly List<string> _allowedFileExtensions = new List<string>();

    public FileAllowedExtensionFilter(string? extensions = null)
    {
        _allowedFileExtensions = !string.IsNullOrWhiteSpace(extensions)
       ? extensions.Split(',').Select(e => e.Trim().ToLowerInvariant()).ToList()
       : Constants.ALLOWED_DOCUMENT_EXTENSIONS
           .Concat(Constants.ALLOWED_IMAGE_EXTENSIONS)
           .Concat(Constants.ALLOWED_VIDEO_EXTENSIONS)
           .Concat(Constants.ALLOWED_OTHER_EXTENSIONS)
           .ToList();
    }

    public async Task OnActionExecutionAsync(ActionExecutingContext context, ActionExecutionDelegate next)
    {
        if (context.ActionArguments.TryGetValue("files", out var value) && value is List<IFormFile> files)
        {
            foreach (var file in files)
            {
                var extension = Path.GetExtension(file.FileName).ToLowerInvariant();
                if (!_allowedFileExtensions.Contains(extension))
                {
                    context.HttpContext.Response.StatusCode = StatusCodes.Status400BadRequest;
                    await context.HttpContext.Response.WriteAsJsonAsync(BaseApiRes<string>.FromError(
                        $"File extension '{extension}' is not allowed. Allowed extensions are: {string.Join(", ", _allowedFileExtensions)}"));
                    return;
                }
            }
        }

        await next();
    }
}
