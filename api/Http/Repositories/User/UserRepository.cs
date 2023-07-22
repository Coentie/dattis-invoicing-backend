using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using api.data;
using api.Entities;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Internal;
using AppUser = api.Entities.User;

namespace api.Http.Repositories.User
{
    public class UserRepository : IRepository
    {
        private readonly DataContext _context;
        
        // Constructor.
        public UserRepository(DataContext context) {
            this._context = context;
        }

        // Creates a new user.
        async public Task<AppUser> Create(AppUser user) {
            _context.Users.Add(user);
            user.Id = await _context.SaveChangesAsync();

            return user;
        }

        // Finds a user by email.
        async public Task<AppUser> FindByEmail(string email) {
            var user = await _context.Users.FirstAsync(u => u.Email == email);
            return user;
        }

        async public Task<AppUser> FindById(int id) {
            AppUser user = await _context.Users.FindAsync(id);
            return user;
        }
    }
}