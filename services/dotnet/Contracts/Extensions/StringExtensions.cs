namespace Contracts.Extensions;

public static class StringExtensions
{
    public static string ToSlug(this string input)
    {
        if (string.IsNullOrWhiteSpace(input))
            return string.Empty;

        // Convert to lower case
        input = input.ToLowerInvariant();

        // Replace spaces and special characters with hyphens
        var slug = System.Text.RegularExpressions.Regex.Replace(input, @"[^\w\s-]", string.Empty);
        slug = System.Text.RegularExpressions.Regex.Replace(slug, @"\s+", "-").Trim('-');

        return slug;
    }
}
