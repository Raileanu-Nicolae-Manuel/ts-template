export const passwordValidation = (password: string) => {
    const specialChars = "!@#$%^&*()_+-=[]{}|;:,.<>?";
    return password.split("").some(char => specialChars.includes(char)) &&
        password.length >= 8 && 
        password.toLowerCase() !== password && 
        password.toUpperCase() !== password;
}

export const emailValidation = (email: string) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}

