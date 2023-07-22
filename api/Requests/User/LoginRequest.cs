using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace api.Requests.User
{
    public class LoginRequest
    {
        public string Password {get; set;}
        public string Email {get; set;}
    }
}