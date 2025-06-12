namespace file_service.Extensions;

public static class FileExtension
{
    public static bool IsVideoFile(this IFormFile file)
    {
        if (file == null || file.Length == 0)
        {
            return false;
        }

        var allowedExtensions = new[] { ".mp4", ".avi", ".mov", ".mkv", ".flv" };
        var fileExtension = Path.GetExtension(file.FileName).ToLowerInvariant();

        return allowedExtensions.Contains(fileExtension);
    }
}