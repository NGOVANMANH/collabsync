public static class Constants
{
    public const long MAX_FILE_SIZE = 100 * ONE_MB; // 100 MB
    public const long ONE_KB = 1024; // 1 KB
    public const long ONE_MB = 1024 * 1024; // 1 MB
    public const long ONE_GB = 1024 * 1024 * 1024; // 1 GB
    public static readonly List<string> FILE_STORAGE_PATHS = new List<string>
    {
        "Uploads",
        "Uploads/Images",
        "Uploads/Videos",
        "Uploads/Other"
    };
    public static readonly List<string> ALLOWED_VIDEO_EXTENSIONS = new List<string>
    {
        ".mp4", ".avi", ".mkv", ".mov", ".flv"
    };
    public static readonly List<string> ALLOWED_IMAGE_EXTENSIONS = new List<string>
    {
        ".jpg", ".jpeg", ".png", ".gif"
    };
    public static readonly List<string> ALLOWED_DOCUMENT_EXTENSIONS = new List<string>
    {
        ".pdf", ".docx", ".txt"
    };
    public static readonly List<string> ALLOWED_OTHER_EXTENSIONS = new List<string>
    {
        ".zip", ".rar", ".7z", ".tar", ".gz", ".exe"
    };
}