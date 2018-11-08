package com.marcoberardelli.securecamhub.repository;

import com.marcoberardelli.securecamhub.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;


@Repository
public interface UserRepository extends JpaRepository<User, Long> {

    public User findByUsername(String username);

}
