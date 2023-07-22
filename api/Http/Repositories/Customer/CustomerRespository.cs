using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using api.data;
using api.Entities;
using api.Http.Repositories;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Internal;
using AppCustomer = api.Entities.Customer;

namespace api.Repositories.Http.Customer
{
    public class CustomerRespository : IRepository
    {
        private readonly DataContext _context;

        public CustomerRespository(DataContext context) {
            this._context = context;
        }

        public async Task<List<AppCustomer>> All() {
            return await this._context.Customers.ToListAsync();
        }

        public async Task<AppCustomer> FindById(int id) {
            return await this._context.Customers.FirstAsync(c => c.Id == id);
        }

        public void Destroy(int id) {
            var customer = this._context.Customers.FirstOrDefault(c => c.Id == id);

            if(customer != null) {
                this._context.Remove(customer);
                this._context.SaveChanges();
            }
        }

        public async Task<AppCustomer> Create(AppCustomer customer) {
            this._context.Customers.Add(customer);
            customer.Id = await this._context.SaveChangesAsync();

            return customer;
        }
    }
}