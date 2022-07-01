import React from 'react';

export default function LoginForm() {
    return (
    <div>
        <h3>Enter login details here </h3>
        <form 
        encType="application/x-www-form-urlencoded"
        action="http://localhost:9000/api/signin"
        method="post"
        >
        <label>
            Name:
        <input type="string" name="Username" />
        </label>
        <label>
            Password:
        <input type="password" name="Password" />
        </label>
        <input type="submit" value="submit" />
        </form>
    </div>
    )
}