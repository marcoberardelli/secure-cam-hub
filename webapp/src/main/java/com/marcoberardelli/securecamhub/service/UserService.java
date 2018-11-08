package com.marcoberardelli.securecamhub.service;

import com.marcoberardelli.securecamhub.model.User;
import com.marcoberardelli.securecamhub.model.UserDTO;
import com.marcoberardelli.securecamhub.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class UserService {


    @Autowired
    PasswordEncoder passwordEncoder;

    @Autowired
    private UserRepository userRepository;


    public void addUser(UserDTO userDTO) {
        userDTO.setPassword(passwordEncoder.encode(userDTO.getPassword()));
        userDTO.setMatchingPassword(passwordEncoder.encode(userDTO.getMatchingPassword()));
        System.out.println(userDTO.getPassword() + "\n" + userDTO.getMatchingPassword());
        User user = new User(userDTO);
        // TODO: add role
        userRepository.save(user);
    }
}
