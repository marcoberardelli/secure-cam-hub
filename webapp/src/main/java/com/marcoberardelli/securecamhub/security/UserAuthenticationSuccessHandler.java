package com.marcoberardelli.securecamhub.security;

import java.io.IOException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.DefaultRedirectStrategy;
import org.springframework.security.web.RedirectStrategy;
import org.springframework.security.web.authentication.AuthenticationSuccessHandler;
import org.springframework.stereotype.Component;


@Component
public class UserAuthenticationSuccessHandler implements AuthenticationSuccessHandler {


    private RedirectStrategy redirectStrategy = new DefaultRedirectStrategy();

    @Override
    public void onAuthenticationSuccess(HttpServletRequest hsr, HttpServletResponse hsr1, Authentication a) throws IOException {
        redirectStrategy.sendRedirect(hsr, hsr1, "/home");
    }

}