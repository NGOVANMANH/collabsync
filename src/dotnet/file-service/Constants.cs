public static class Constants
{
    public const long MAX_FILE_SIZE = 100 * MB; // 100 MB
    public const long KB = 1024; // 1 KB
    public const long MB = 1024 * 1024; // 1 MB
    public const long GB = 1024 * 1024 * 1024; // 1 GB
    public static readonly HashSet<string> FILE_STORAGE_PATHS = new()
    {
        "Uploads",
        "Uploads/Images",
        "Uploads/Videos",
        "Uploads/Audios",
        "Uploads/Documents",
        "Uploads/Other"
    };
}