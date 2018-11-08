package com.marcoberardelli.securecamhub.controller;


import com.marcoberardelli.securecamhub.model.UserDTO;
import com.marcoberardelli.securecamhub.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;

import javax.validation.Valid;

@Controller
public class MvcController {


    @Autowired
    private UserService userService;

    @GetMapping(value = "/")
    public String indexPage() {
        return "index";
    }

    @GetMapping(value = "/home")
    public String homePage() { return "home"; }

    @GetMapping(value = "/register")
    public String registerPage(Model model) {
        model.addAttribute("userDTO", new UserDTO());
        return "register";
    }

    @PostMapping(value = "/register")
    public void registerUser(@Valid @ModelAttribute UserDTO userDTO) {

        userService.addUser(userDTO);
    }
}
