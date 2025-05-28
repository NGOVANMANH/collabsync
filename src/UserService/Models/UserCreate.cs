using System.ComponentModel.DataAnnotations;

namespace UserService.Models;

public class UserCreate
{
    [Required]
    public string UserName { get; set; } = null!;
    [Required]
    [EmailAddress]
    public string Email { get; set; } = null!;
    [Required]
    public string Password { get; set; } = null!;
    public DateTime DateOfBirth { get; set; }
}
