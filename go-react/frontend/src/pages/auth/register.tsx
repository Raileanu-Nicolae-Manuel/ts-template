import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { ApiService } from "@/service/api";
import { emailValidation, passwordValidation } from "@/service/util/input.validation";
import { useState } from "react";
import { NavLink } from "react-router-dom";

export default function RegisterPage(){
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const usersService = new ApiService().usersService;

  const handleRegister = async () => {
    const response = await usersService.register({email, username, password});
    console.log(response);
  }

  const handleValidInput = () => {
    return passwordValidation(password) && 
    emailValidation(email) && 
    username.length > 3 && 
    confirmPassword === password;
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { id, value } = e.target;
    console.log(id, value);
    if (id === "email") setEmail(value);
    else if (id === "username") setUsername(value);
    else if (id === "password") setPassword(value);
    else if (id === "confirm-password") setConfirmPassword(value);
  }

  return (
    <div className="size-full grid place-items-center">
      <Card className="w-[60%] max-w-[22rem] h-fit">
        <CardHeader>
          <CardTitle>Register</CardTitle>
        </CardHeader>
        <CardContent className="pb-4">
          <form>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Input id="username" placeholder="Username" onChange={handleChange} />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Input id="email" placeholder="Email" onChange={handleChange} />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Input id="password" name="password" type="password" placeholder="Password" onChange={handleChange}   />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Input id="confirm-password" type="password" placeholder="Confirm Password" onChange={handleChange} />
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex flex-col justify-between gap-4">
          <Button className="dark:bg-slate-200 w-full" disabled={!handleValidInput()} onClick={handleRegister}>Register</Button>
          <div className="dark:text-slate-300 flex-1">
            <span>
              Do you have an account?
            </span>
            <NavLink to='/login' className='underline ml-1'>Login</NavLink>
          </div>
        </CardFooter>
      </Card>
    </div>
  )
}