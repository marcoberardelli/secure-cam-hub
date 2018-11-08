package com.marcoberardelli.securecamhub.model;


import java.io.Serializable;
import javax.persistence.*;

@Entity
@Table(name = "user")
public class User implements Serializable {

    @Id
    @Column
    private long id;

    @Column
    private String username;

    @Column
    private String password;


    public User () {}

    public User (UserDTO userDTO) {
        this.username = userDTO.getEmail();
        this.password = userDTO.getPassword();
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getUsername() { return username; }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() { return password; }

    public void setPassword(String password) {
        this.password = password;
    }



    // Code for using roles
    /*
    @ManyToMany(cascade = CascadeType.ALL, fetch = FetchType.EAGER)
    @JoinTable(
            name = "users_roles",
            joinColumns = {@JoinColumn(name = "user_id")},
            inverseJoinColumns = {@JoinColumn(name = "role_id")})
    @JsonManagedReference
    private Set<Role> roles = new HashSet<>();

    //Getter and setters
    public Set<Role> getRoles() {
        return roles;
    }

    public void setRoles(Set<Role> roles) {
        this.roles = roles;
    }
    */
}
