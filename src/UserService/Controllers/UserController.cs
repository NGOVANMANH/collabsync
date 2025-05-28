using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using UserService.Data;
using UserService.Models;

namespace UserService.Controllers
{
    [Route("api/[controller]s")]
    [ApiController]
    public class UserController : ControllerBase
    {
        private readonly UserServiceDbContext _context;

        public UserController(UserServiceDbContext context)
        {
            _context = context;
        }

        [HttpGet]
        public async Task<IActionResult> GetUsers()
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would retrieve users from a database or another service.
            return Ok(await _context.Users.ToListAsync());
        }
        [HttpGet("{id}")]
        public async Task<IActionResult> GetUserById(Guid id)
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would retrieve a specific user by their ID from a database or another service.

            return Ok(await _context.Users.FirstOrDefaultAsync(u => u.Id == id));
        }
        [HttpPost]
        public async Task<IActionResult> CreateUser([FromBody] UserCreate user)
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would save the new user to a database or another service.

            var createdUser = await _context.Users.AddAsync(new Models.User
            {
                Id = Guid.NewGuid(),
                UserName = user.UserName,
                Email = user.Email,
                DateOfBirth = user.DateOfBirth
            });

            await _context.SaveChangesAsync();

            return CreatedAtAction(nameof(GetUserById), new { id = createdUser.Entity.Id }, createdUser.Entity);
        }
        [HttpPost("fake-users")]
        public async Task<IActionResult> FakeUsers()
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would create fake users for testing purposes.

            var users = new List<User>();

            for (int i = 0; i < 10; i++)
            {
                // Here you would typically add the user to the database context and save changes.
                // _context.Users.Add(user);
                users.Add(new User
                {
                    Id = Guid.NewGuid(),
                    UserName = $"Fake User {i + 1}",
                    Email = $"fakeuser{i + 1}@gmail.com",
                    DateOfBirth = DateTime.UtcNow.AddYears(-20 - i), // Example: 20 years old minus i
                });
            }

            await _context.AddRangeAsync(users);
            await _context.SaveChangesAsync();

            return Ok(users);
        }

        [HttpDelete]
        public async Task<IActionResult> DeleteUsers()
        {
            // This is a placeholder for the actual implementation.
            // In a real application, you would delete the specified user from a database or another service.
            _context.Users.RemoveRange(_context.Users);
            await _context.SaveChangesAsync();
            return Ok(new { Message = "User deleted successfully" });
        }
    }
}
