using api.data;
using api.Http.Requests.Customer;
using api.Repositories.Http.Customer;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using AppCustomer = api.Entities.Customer;

namespace api.Controllers.Customer
{
    [ApiController]
    [Authorize]
    [Route("api/customers")]
    public class CustomerController : ControllerBase
    {
            private CustomerRespository _repository;

            public CustomerController(CustomerRespository respository) {
                this._repository = respository;
            }

            [HttpGet]
            public async Task<ActionResult<IEnumerable<AppCustomer>>> Index() {
                var customers = await this._repository.All();

                return Ok(customers);
            }

            [HttpPost]
            public async Task<ActionResult<AppCustomer>> Store([FromBody] CustomerRequest request) {
                var customer = await this._repository.Create( new AppCustomer {
                    Name = request.Name,
                    Email = request.Email,
                });

                return Ok(customer);
            }

            [HttpDelete]
            [Route("{id}")]
            public IActionResult Destroy(int id) {
                _repository.Destroy(id);

                return Ok();
            }

            [HttpGet()]
            [Route("{id}")]
            public async Task<IActionResult> Get(int id) {
                var customer = await _repository.FindById(id);

                if(customer == null) {
                    return NotFound();
                }

                return Ok(customer);
            }
    }
}