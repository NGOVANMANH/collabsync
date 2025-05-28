namespace Contracts.Base.Models;

public interface IBaseEntity
{
    /// <summary>
    /// Gets or sets the unique identifier for the entity.
    /// </summary>
    Guid Id { get; set; }

    /// <summary>
    /// Gets or sets the creation date of the entity.
    /// </summary>
    DateTime CreatedAt { get; set; }

    /// <summary>
    /// Gets or sets the last modification date of the entity.
    /// </summary>
    DateTime? ModifiedAt { get; set; }
}
