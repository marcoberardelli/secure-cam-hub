package com.marcoberardelli.securecamhub.model;


import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Size;


public class UserDTO {

    @NotBlank(message = "Please enter a valid email")
    @Email
    private String email;

    @NotBlank
    @Size(min = 2, max = 32 , message = "Password must be between 8 and 32 characters")
    private String password;

    @NotBlank
    private String matchingPassword;

    public UserDTO() {}

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getMatchingPassword() {
        return matchingPassword;
    }

    public void setMatchingPassword(String matchingPassword) {
        this.matchingPassword = matchingPassword;
    }
}
