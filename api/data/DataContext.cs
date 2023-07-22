using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using api.Entities;
using Microsoft.EntityFrameworkCore;

namespace api.data
{
    public class DataContext : DbContext
    {
            public DataContext(DbContextOptions options): base(options) {

            }

            public DbSet<User> Users {get ; set; }
            public DbSet<Customer> Customers {get; set;} 
    }
}