namespace file_service.Utils;

public static class FileStorage
{
    public static void EnsureDirectoryExists(string directoryPath)
    {
        if (!Directory.Exists(directoryPath))
        {
            Directory.CreateDirectory(directoryPath);
        }
    }

    public static string GetLocalFilePath(string fileUrl)
    {
        return Path.Combine(Directory.GetCurrentDirectory(), Constants.FILE_STORAGE_ROOT, fileUrl);
    }
}