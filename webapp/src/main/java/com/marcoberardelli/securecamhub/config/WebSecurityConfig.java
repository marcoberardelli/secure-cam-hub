package com.marcoberardelli.securecamhub.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

@Configuration
@EnableWebSecurity
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {



    @Autowired
    private UserDetailsService customUserDetailsService;

    @Autowired
    public void configureAuth(AuthenticationManagerBuilder auth) throws Exception {
        auth.userDetailsService(customUserDetailsService).passwordEncoder(passwordEncoder());
    }

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
                .authorizeRequests()
                //The credentials are "utente" and "1234", but the pass is crypted.
                //Just put this in the db:
                //INSERT INTO `users` (`id`, `name`, `password`) VALUES(11, 'utente', '$2a$10$gpUGDkA3uhtT/BM1WMrt0ek/mZQXC5e4V9zPRBXrjd1AVeixL8KQO');
                // IN THE DB THE STRING SHOULD BE "ROLE_USER"
                .antMatchers("/register").permitAll()
                .antMatchers("/css/**").permitAll()
                .antMatchers("/").permitAll()
                .anyRequest().authenticated()
                .and()
                .formLogin().permitAll()
                //.loginPage("/login")
                .and()
                .logout().permitAll()
                .and().csrf().disable();
    }

    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }
}
