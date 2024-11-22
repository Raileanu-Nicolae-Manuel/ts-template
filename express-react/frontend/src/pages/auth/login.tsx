import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { NavLink } from "react-router-dom"


export default function LoginPage(){
  return (
    <div className="size-full grid place-items-center">
      <Card className="w-[60%] max-w-[22rem] h-fit">
        <CardHeader>
          <CardTitle>Login</CardTitle>
        </CardHeader>
        <CardContent className="pb-4">
          <form>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                {/* <Label htmlFor="email">Email</Label> */}
                <Input id="email" placeholder="Email" />
              </div>
              <div className="flex flex-col space-y-1.5">
                {/* <Label htmlFor="password">Password</Label> */}
                <Input id="password" type="password" placeholder="Password" />
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex flex-col justify-between gap-4">
          <Button className="dark:bg-slate-200 w-full">Login</Button>
          <div className="dark:text-slate-300 flex-1">
            <span>
              Don't have an account?
            </span>
            <NavLink to='/register' className='underline ml-1'>Register</NavLink>
          </div>
        </CardFooter>
      </Card>
    </div>
  )
}