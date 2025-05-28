using System.Threading.Tasks;
using BlogService.Data;
using BlogService.Models;
using Microsoft.AspNetCore.Mvc;

namespace BlogService.Controllers
{
    [Route("api/[controller]s")]
    [ApiController]
    public class BlogController : ControllerBase
    {
        private readonly BlogServiceDbContext _context;

        public BlogController(BlogServiceDbContext context)
        {
            _context = context;
        }

        [HttpGet]
        public IActionResult GetBlogs()
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would retrieve blogs from a database or another service.
            return Ok(new { Message = "List of blogs" });
        }

        [HttpGet("{id}")]
        public IActionResult GetBlogById(int id)
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would retrieve a specific blog by its ID from a database or another service.
            return Ok(new { Message = $"Details of blog with ID {id}" });
        }

        [HttpPost]
        public IActionResult CreateBlog([FromBody] object blog)
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would save the new blog to a database or another service.
            return CreatedAtAction(nameof(GetBlogById), new { id = 1 }, blog);
        }

        [HttpPost("fake-blogs")]
        public async Task<IActionResult> FakeBlogs()
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would create fake blogs for testing purposes.

            var blogs = new List<Blog>();

            for (int i = 0; i < 10; i++)
            {
                // Here you would typically add the blog to the database context and save changes.
                // _context.Blogs.Add(blog);
                blogs.Add(new Blog
                {
                    Id = Guid.NewGuid(),
                    Title = $"Fake Blog Title {i + 1}",
                    Summary = "This is a fake blog summary.",
                    Content = "This is the content of the fake blog. It is meant for testing purposes only.",
                    AuthorId = Guid.NewGuid(),
                });
            }

            await _context.Blogs.AddRangeAsync(blogs);
            await _context.SaveChangesAsync();

            return Ok(new { Message = "Fake blogs created" });
        }

        [HttpDelete]
        public IActionResult DeleteAllBlogs()
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would delete all blogs from the database or another service.

            _context.Blogs.RemoveRange(_context.Blogs);
            _context.SaveChanges();

            return Ok(new { Message = "All blogs deleted" });
        }
    }
}
