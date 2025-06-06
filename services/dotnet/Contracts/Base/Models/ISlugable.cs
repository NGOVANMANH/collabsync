namespace Contracts.Base.Models;

public interface ISlugable
{
    /// <summary>
    /// Gets or sets the slug for the entity.
    /// </summary>
    string Slug { get; set; }

    /// <summary>
    /// Gets or sets the title of the entity.
    /// </summary>
    string Title { get; set; }
}
