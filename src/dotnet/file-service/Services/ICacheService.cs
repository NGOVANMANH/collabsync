
using Microsoft.Extensions.Caching.Distributed;

namespace file_service.Services;

public interface ICacheService
{
    Task<T?> GetAsync<T>(string key);
    Task SetAsync<T>(string key, T value, DistributedCacheEntryOptions options);
    Task RemoveAsync(string key);
}
