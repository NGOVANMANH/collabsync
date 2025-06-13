
namespace file_service.Models.API;

public class BaseApiRes<T>
{
    public bool Success { get; set; }
    public string? Message { get; set; }
    public T? Data { get; set; }
    public BaseApiRes(bool success, string? message = null, T? data = default)
    {
        Success = success;
        Message = message;
        Data = data;
    }

    public static BaseApiRes<T> FromSuccess(T data, string? message = null)
    {
        return new BaseApiRes<T>(true, message, data);
    }

    public static BaseApiRes<T> FromError(string message, T? data = default)
    {
        return new BaseApiRes<T>(false, message, data);
    }
    public static BaseApiRes<T> FromException(Exception ex, T? data = default)
    {
        return new BaseApiRes<T>(false, ex.Message, data);
    }
}