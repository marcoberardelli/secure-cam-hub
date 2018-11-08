package com.marcoberardelli.securecamhub.security;


import java.util.Collection;
import java.util.HashSet;
import java.util.Set;

import com.marcoberardelli.securecamhub.model.Role;
import com.marcoberardelli.securecamhub.model.User;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

public class CustomUserDetails implements UserDetails{


    private User user;

    public CustomUserDetails(User user) {
        this.user = user;
    }


    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        Collection<GrantedAuthority> authorities = new HashSet<>();
        authorities.add(new SimpleGrantedAuthority("USER"));

        return authorities;


        // Code for using roles
        /*
        Set<Role> roles = user.getRoles();
        roles.forEach((role) -> {
            authorities.add(new SimpleGrantedAuthority(role.getRole()));
        });
        return authorities;
        */
    }

    @Override
    public String getUsername() {
        return user.getUsername();
    }

    @Override
    public boolean isAccountNonExpired() {
        return true;
    }

    @Override
    public boolean isAccountNonLocked() {
        return true;
    }

    @Override
    public boolean isCredentialsNonExpired() {
        return true;
    }

    @Override
    public boolean isEnabled() {
        return true;
    }

    @Override
    public String getPassword() {
        return user.getPassword();
    }

}