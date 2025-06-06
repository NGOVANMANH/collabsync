using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using Contracts.Base.Models;
using Microsoft.EntityFrameworkCore;
using MongoDB.Bson.Serialization.Attributes;

namespace BlogService.Models;

[Index(nameof(Slug), IsUnique = true)]
public class Blog : IBaseEntity, ISlugable, ISoftDelete
{
    [BsonId]
    public Guid Id { get; set; }

    [Required]
    [MaxLength(100)]
    [MinLength(3)]
    public string Title { get; set; } = null!;

    [Required]
    [MaxLength(500)]
    [MinLength(10)]
    public string Slug { get; set; } = null!;

    [Required]
    [MaxLength(5000)]
    [MinLength(20)]
    public string Summary { get; set; } = null!;

    [Required]
    [MaxLength(10000)]
    [MinLength(50)]
    [DataType(DataType.Html)]
    public string Content { get; set; } = null!;

    public Guid AuthorId { get; set; }

    [ForeignKey(nameof(AuthorId))]
    public User Author { get; set; } = null!;

    public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
    public DateTime? ModifiedAt { get; set; } = null;
    public bool IsDeleted { get; set; }
    public DateTime? DeletedAt { get; set; }
}
