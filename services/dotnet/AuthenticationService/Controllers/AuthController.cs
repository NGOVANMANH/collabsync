using Microsoft.AspNetCore.Identity.Data;
using Microsoft.AspNetCore.Mvc;

namespace AuthenticationService.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class AuthController : ControllerBase
    {
        [HttpPost]
        public IActionResult Login([FromBody] LoginRequest request)
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would validate the user credentials and return a token or user information.
            if (request.Email == "test" && request.Password == "password")
            {
                return Ok(new { Token = "fake-jwt-token" });
            }
            return Unauthorized();
        }

        [HttpPost("register")]
        public IActionResult Register([FromBody] RegisterRequest request)
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would save the new user to a database and return user information.
            if (ModelState.IsValid)
            {
                return CreatedAtAction(nameof(Register), new { id = 1 }, request);
            }
            return BadRequest(ModelState);
        }
    }
}
