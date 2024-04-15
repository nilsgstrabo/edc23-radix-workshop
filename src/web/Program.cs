using Azure.Identity;
using Microsoft.Extensions.Azure;

namespace edc23_radix_workshop_1;

public class Program
{
    public static void Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);

        // Add services to the container.
        builder.Services.AddRazorPages();

        // builder.Services.AddAzureClients(builder=> {
        //     builder.AddBlobServiceClient(new Uri("https://edc23radixworkshop1.blob.core.windows.net"));

        //     // Use DefaultAzureCredential when debugging locally, otherwise WorkloadIdentityCredential
        //     builder.UseCredential((provider) => provider.GetService<IWebHostEnvironment>()!.IsDevelopment() ? new DefaultAzureCredential() : new WorkloadIdentityCredential());
        // });

        var app = builder.Build();

        // Configure the HTTP request pipeline.
        if (!app.Environment.IsDevelopment())
        {
            app.UseExceptionHandler("/Error");
        }
        app.UseStaticFiles();

        app.UseRouting();

        app.UseAuthorization();

        app.MapRazorPages();
        
        app.MapControllers();

        app.Run();
    }
}
