using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using api.Http.Repositories.User;
using api.Http.Resources;
using api.Repositories;
using api.Requests.User;
using api.Services.JWT;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http.HttpResults;
using Microsoft.AspNetCore.Mvc;
using AppUser = api.Entities.User;

namespace api.Controllers
{
    [Route("api/auth")]
    [ApiController]
    public class AuthController : ControllerBase
    {
        private UserRepository _repository;

        private ITokenService _jwtService;

        // Auth controller.
        public AuthController(UserRepository repository, ITokenService jwtService) {
            this._repository = repository;
            this._jwtService = jwtService;
        }
        
        [AllowAnonymous]
        [HttpPost("login")]
        public async Task<ActionResult<UserResource>> login(LoginRequest request) {
            var user = await this._repository.FindByEmail(request.Email);

            if(user == null) {
                return BadRequest("User/Password combination is not found");
            }

            if(! BCrypt.Net.BCrypt.Verify(request.Password, user.Password)) {
                return BadRequest("User/Password combination is not found");
            }

            var jwtToken = this._jwtService.CreateToken(user);

            Response.Cookies.Append("JWT", jwtToken, new CookieOptions{
                HttpOnly = true
            });

            return new UserResource {
                UserName = user.UserName,
                Email = user.Email,
                Token = jwtToken
            };             
        }

        // [AllowAnonymous]
        // [HttpPost("refresh")]
        // public async Task<ActionResult<UserResource>> refresh() {
        //     int userId;

        //     try {
        //         var cookieToken = Request.Cookies["jwt"];
        //         var token = this._jwtService.Verify(cookieToken);
        //         userId = int.Parse(token.Issuer);
                
        //         Response.Cookies.Delete("JWT");
        //     }catch(Exception e) {
        //         return Unauthorized();
        //     }

        //     var user = await this._repository.FindById(userId);
        //     var jwtToken = this._jwtService.CreateToken(user);

        //     Response.Cookies.Append("JWT", jwtToken, new CookieOptions{
        //         HttpOnly = true
        //     });

        //     return new UserResource {
        //         UserName = user.UserName,
        //         Email = user.Email,
        //         Token = _jwtService.CreateToken(user)
        //     };
        // }

        [AllowAnonymous]
        [HttpPost("register")]
        public async Task<UserResource> Register(RegisterRequest request) {
            var user = new AppUser 
            {
                UserName = request.Name, 
                Email = request.Email,
                Password = BCrypt.Net.BCrypt.HashPassword(request.Password)
            };

            user = await this._repository.Create(user);

            return new UserResource {
                UserName = user.UserName,
                Email = user.Email,
                Token = _jwtService.CreateToken(user)
            };
        }
    }
}