using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using api.data;
using api.Http.Repositories.User;
using api.Repositories.Http.Customer;
using api.Services.JWT;
using Microsoft.EntityFrameworkCore;

namespace api.Extensions
{
    static public class ApplicationServiceExtensions
    {
        public static IServiceCollection AddApplicationServices(this IServiceCollection services, IConfiguration config) {
            services.AddControllers();
            services.AddEndpointsApiExplorer();

            services.AddScoped<UserRepository>();
            services.AddScoped<CustomerRespository>();

            services.AddScoped<ITokenService, TokenService>();
            services.AddSwaggerGen();
            services.AddDbContext<DataContext>(opt => {
                opt.UseMySql(
                    config.GetConnectionString("DefaultConnection"), new MySqlServerVersion(new Version(8, 0, 29))
                );
            });
            services.AddCors();

            return services;
        }
    }
}