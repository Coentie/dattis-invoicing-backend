using System;
using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace api.Entities
{
	public class User
	{
		[Key]
		public int Id { get; set; }
		
		public string UserName { get; set; }

		[JsonIgnore]
		public string Password {get; set; }

		public string Email {get; set;}

		public User()
		{

		}
	}
}

