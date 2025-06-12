using file_service.Models.API;
using Microsoft.AspNetCore.Mvc.Filters;

namespace file_service.Filters;

public class FileTypeAllowedFilter : Attribute, IAsyncActionFilter
{
    private readonly List<string> _allowedFileTypes = new List<string>();

    public FileTypeAllowedFilter(string? fileTypeCsv = null)
    {
        if (string.IsNullOrEmpty(fileTypeCsv))
        {
            _allowedFileTypes.AddRange(new[] { "image/jpeg", "image/png", "application/pdf" });
        }
        else
        {
            _allowedFileTypes.AddRange(fileTypeCsv.Split(',').Select(ft => ft.Trim().ToLowerInvariant()));
        }
    }

    public async Task OnActionExecutionAsync(ActionExecutingContext context, ActionExecutionDelegate next)
    {
        if (context.ActionArguments.TryGetValue("files", out var value) && value is List<IFormFile> files)
        {
            foreach (var file in files)
            {
                var fileContentType = file.ContentType.ToLowerInvariant();
                if (!_allowedFileTypes.Contains(fileContentType.Split('/')[0]))
                {
                    context.HttpContext.Response.StatusCode = StatusCodes.Status400BadRequest;
                    await context.HttpContext.Response.WriteAsJsonAsync(BaseApiRes<string>.FromError(
                        $"File types not support: {string.Join(", ", _allowedFileTypes)}"));
                    return;
                }
            }
        }

        await next();
    }
}
