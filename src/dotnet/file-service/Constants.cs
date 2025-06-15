namespace file_service;

public static class Constants
{
    // -----------------------
    // Storage Size Limits
    // -----------------------
    public const long KB = 1024;
    public const long MB = 1024 * KB;
    public const long GB = 1024 * MB;

    public const long IMAGE_MAX_SIZE = 5 * MB;         // 5 MB
    public const long VIDEO_MAX_SIZE = 55 * MB;        // 55 MB
    public const long AUDIO_MAX_SIZE = 20 * MB;        // 20 MB
    public const long DOCUMENT_MAX_SIZE = 10 * MB;     // 10 MB
    public const long OTHER_MAX_SIZE = 10 * MB;        // 10 MB

    // -----------------------
    // Storage Paths
    // -----------------------
    public const string FILE_STORAGE_ROOT = "Uploads";
    public const string FILE_STORAGE_IMAGES = $"{FILE_STORAGE_ROOT}/Images";
    public const string FILE_STORAGE_VIDEOS = $"{FILE_STORAGE_ROOT}/Videos";
    public const string FILE_STORAGE_AUDIOS = $"{FILE_STORAGE_ROOT}/Audios";
    public const string FILE_STORAGE_DOCUMENTS = $"{FILE_STORAGE_ROOT}/Documents";
    public const string FILE_STORAGE_OTHER = $"{FILE_STORAGE_ROOT}/Other";
    public const string FILE_STORAGE_FFMPEG_FOLDER = $"{FILE_STORAGE_VIDEOS}/ffmpeg";

    public static readonly HashSet<string> FILE_STORAGE_PATHS = new()
    {
        FILE_STORAGE_ROOT,
        FILE_STORAGE_IMAGES,
        FILE_STORAGE_VIDEOS,
        FILE_STORAGE_FFMPEG_FOLDER,
        FILE_STORAGE_AUDIOS,
        FILE_STORAGE_DOCUMENTS,
        FILE_STORAGE_OTHER
    };

    // -----------------------
    // Allowed MIME Types
    // -----------------------
    public static readonly HashSet<string> ALLOWED_IMAGE_FILE_TYPES = new()
    {
        "image/jpeg",
        "image/png",
        "image/gif",
        "image/avif",
        "image/bmp",
        "image/vnd.microsoft.icon"
    };

    public static readonly HashSet<string> ALLOWED_VIDEO_FILE_TYPES = new()
    {
        "video/mp4",
        "video/mpeg",
        "video/mp2t",
        "video/webm"
    };

    public static readonly HashSet<string> ALLOWED_AUDIO_FILE_TYPES = new()
    {
        "audio/mpeg",
        "audio/wav",
        "audio/ogg"
    };

    public static readonly HashSet<string> ALLOWED_DOCUMENT_FILE_TYPES = new()
    {
        "text/plain",
        "application/pdf",
        "application/msword",
        "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        "application/vnd.ms-powerpoint",
        "application/vnd.openxmlformats-officedocument.presentationml.presentation"
    };

    public static readonly HashSet<string> ALLOWED_OTHER_FILE_TYPES = new()
    {
        "application/octet-stream",
        "application/zip",
        "application/x-gzip"
    };

    public static readonly HashSet<string> ALLOWED_FILE_TYPES =
        ALLOWED_IMAGE_FILE_TYPES
        .Union(ALLOWED_VIDEO_FILE_TYPES)
        .Union(ALLOWED_AUDIO_FILE_TYPES)
        .Union(ALLOWED_DOCUMENT_FILE_TYPES)
        .Union(ALLOWED_OTHER_FILE_TYPES)
        .ToHashSet();
}
