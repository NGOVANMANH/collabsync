using Microsoft.AspNetCore.Mvc;

namespace BlogService.Controllers
{
    [Route("api/[controller]s")]
    [ApiController]
    public class BlogController : ControllerBase
    {
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
    }
}
