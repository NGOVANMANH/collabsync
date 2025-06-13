namespace file_service.Extensions;

public static class FileExtension
{
    public static bool IsFileOverSize(this IFormFile file)
    {
        var contentType = file.ContentType?.Trim().ToLowerInvariant();

        if (string.IsNullOrWhiteSpace(contentType))
            return true; // or false, depending on whether you treat unknown types as invalid/oversized

        var maxSize = contentType switch
        {
            _ when Constants.ALLOWED_IMAGE_FILE_TYPES.Contains(contentType) => Constants.IMAGE_MAX_SIZE,
            _ when Constants.ALLOWED_VIDEO_FILE_TYPES.Contains(contentType) => Constants.VIDEO_MAX_SIZE,
            _ when Constants.ALLOWED_AUDIO_FILE_TYPES.Contains(contentType) => Constants.AUDIO_MAX_SIZE,
            _ when Constants.ALLOWED_DOCUMENT_FILE_TYPES.Contains(contentType) => Constants.DOCUMENT_MAX_SIZE,
            _ => Constants.OTHER_MAX_SIZE
        };

        return file.Length > maxSize;
    }
}