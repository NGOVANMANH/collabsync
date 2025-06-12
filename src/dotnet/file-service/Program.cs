using file_service.Filters;
using file_service.Services;
using file_service.Utils;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();

// Check if folder exists, if not create it

foreach (var path in Constants.FILE_STORAGE_PATHS)
{
    var fullPath = Path.Combine(Directory.GetCurrentDirectory(), path);
    FileStorage.EnsureDirectoryExists(fullPath);
}

builder.Services.AddScoped<FileAllowedExtensionFilter>();
builder.Services.AddScoped<IFileService>();

builder.Services.AddStackExchangeRedisCache(options =>
{

    var redisConfig = builder.Configuration
        .GetSection("Redis")
        .Get<file_service.Config.RedisConfig>()
        ?? throw new InvalidOperationException("Redis configuration is missing or invalid.");

    options.Configuration = redisConfig.ConnectionString;
    options.InstanceName = redisConfig.InstanceName;
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
