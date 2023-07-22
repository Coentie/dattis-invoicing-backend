using System.Text;
using api.data;
using api.Extensions;
using api.Repositories;
using api.Services.JWT;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddApplicationServices(builder.Configuration);

// Authentication
builder.Services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
    .AddJwtBearer(options => {
        options.TokenValidationParameters = new TokenValidationParameters
        {
            ValidateIssuerSigningKey = true,
            IssuerSigningKey = new SymmetricSecurityKey(
                Encoding.UTF8.GetBytes(builder.Configuration["TokenKey"])
                ),
            ValidateIssuer = false,
            ValidateAudience = false,
        };
    });
var app = builder.Build();

app.UseHttpsRedirection();

app.UseCors(builder => 
    builder.AllowAnyHeader()
    .AllowAnyMethod()
    .AllowCredentials()
    .WithOrigins(new [] {"http://127.0.0.1:5173", "http://localhost:5173"})
    .WithOrigins("http://127.0.0.1:5173", "http://localhost:5173")
);

app.UseAuthentication();
app.UseAuthorization();

app.MapControllers();

app.Run();

