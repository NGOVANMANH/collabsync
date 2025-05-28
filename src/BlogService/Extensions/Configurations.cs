using BlogService.Data;
using Microsoft.EntityFrameworkCore;
using MongoDB.Driver;

namespace BlogService.Extensions;

public static class Configurations
{
    public static IServiceCollection AddMongoDb(this IServiceCollection services, string uri)
    {
        if (string.IsNullOrWhiteSpace(uri))
        {
            throw new ArgumentException("MongoDB URI cannot be null or empty.", nameof(uri));
        }

        var mongoClient = new MongoClient(uri);

        services.AddDbContext<BlogServiceDbContext>(options =>
        {
            options.UseMongoDB(mongoClient, "blog_service_db");
        });

        return services;
    }
}
