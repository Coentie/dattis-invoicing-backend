using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using api.Entities;


namespace api.Services.JWT
{
    public interface ITokenService
    {
         string CreateToken(User user);

    }
}