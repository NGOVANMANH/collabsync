using file_service.Models.API;
using Microsoft.AspNetCore.Mvc.Filters;

namespace file_service.Filters;

public class FileTypeAllowedFilter : Attribute, IAsyncActionFilter
{
    private readonly HashSet<string> _allowedFileTypes = new HashSet<string>();

    public FileTypeAllowedFilter(string? allowedFileTypes = null)
    {
        if (!string.IsNullOrEmpty(allowedFileTypes))
        {
            _allowedFileTypes = allowedFileTypes.Split(',').Select(t => t.Trim().ToLowerInvariant()).ToHashSet();
        }
        else
        {
            _allowedFileTypes = Constants.ALLOWED_FILE_TYPES;
        }
    }
    public async Task OnActionExecutionAsync(ActionExecutingContext context, ActionExecutionDelegate next)
    {
        if (context.ActionArguments.TryGetValue("files", out var value) && value is List<IFormFile> files)
        {
            foreach (var file in files)
            {
                var fileContentType = file.ContentType.ToLowerInvariant();
                if (!_allowedFileTypes.Contains(fileContentType.ToLowerInvariant()))
                {
                    context.HttpContext.Response.StatusCode = StatusCodes.Status400BadRequest;
                    await context.HttpContext.Response.WriteAsJsonAsync(BaseApiRes<string>.FromError(
                        $"File types not support: [{string.Join(", ", _allowedFileTypes)}]"));
                    return;
                }
            }
        }

        await next();
    }
}
