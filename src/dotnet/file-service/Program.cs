using file_service;
using file_service.Data;
using file_service.Filters;
using file_service.Repositories;
using file_service.Services;
using file_service.Utils;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

DotNetEnv.Env.Load();


builder.Services.AddControllers();
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();

// Check if folder exists, if not create it

foreach (var path in Constants.FILE_STORAGE_PATHS)
{
    var fullPath = Path.Combine(Directory.GetCurrentDirectory(), path);
    FileStorage.EnsureDirectoryExists(fullPath);
}

builder.Services.AddScoped<FileTypeAllowedFilter>();
builder.Services.AddScoped<IFileService, FileService>();

builder.Services.AddDbContext<FileDbContext>(options =>
{
    options.UseNpgsql($"Host={Environment.GetEnvironmentVariable("POSTGRES_HOST")};" +
        $"Port={Environment.GetEnvironmentVariable("POSTGRES_PORT")};" +
        $"Database={Environment.GetEnvironmentVariable("POSTGRES_DB")};" +
        $"Username={Environment.GetEnvironmentVariable("POSTGRES_USER")};" +
        $"Password={Environment.GetEnvironmentVariable("POSTGRES_PASSWORD")}");
});

builder.Services.AddCors(options =>
{
    options.AddPolicy("AllowAllOrigins",
        builder => builder.AllowAnyOrigin()
                          .AllowAnyMethod()
                          .AllowAnyHeader());
});

builder.Services.AddScoped<IFileRepository, FileRepository>();

builder.WebHost.ConfigureKestrel(options =>
{
    options.Limits.MaxRequestBodySize = 60 * Constants.MB; // Set max request body size to 50 MB
});

// builder.Services.AddStackExchangeRedisCache(options =>
// {

//     var redisConfig = builder.Configuration
//         .GetSection("Redis")
//         .Get<file_service.Config.RedisConfig>()
//         ?? throw new InvalidOperationException("Redis configuration is missing or invalid.");

//     options.Configuration = redisConfig.ConnectionString;
//     options.InstanceName = redisConfig.InstanceName;
// });

var app = builder.Build();

app.UseCors("AllowAllOrigins");

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
