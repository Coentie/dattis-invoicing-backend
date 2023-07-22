using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using api.Http.Repositories.User;
using Microsoft.AspNetCore.Mvc;

namespace api.Controllers.User
{
    [Route("users")]
    [ApiController]
    public class UserController : ControllerBase
    {
        private readonly UserRepository _repository;

        public UserController(UserRepository repository) {
            this._repository = repository;
        }
    }
}