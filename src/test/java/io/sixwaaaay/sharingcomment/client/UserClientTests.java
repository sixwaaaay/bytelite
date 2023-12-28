/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.client;

import io.sixwaaaay.sharingcomment.tools.JwtUtil;
import io.sixwaaaay.sharingcomment.transmission.GetMultipleUserReq;
import io.sixwaaaay.sharingcomment.transmission.GetUserReq;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.Set;

@SpringBootTest
public class UserClientTests {
    @Autowired
    private UserClient userClient;
    @Autowired
    private JwtUtil jwtUtil;

    @Test
    void testGetUser() {
        var token = jwtUtil.generateToken("john", "111");

        var userReply = userClient.getUser(new GetUserReq(1L), token);
        Assertions.assertNotNull(userReply);
    }

    @Test
    void testGetUsers() {
        var token = jwtUtil.generateToken("john", "111");
        var usersReply = userClient.getManyUser(new GetMultipleUserReq(Set.of(457232417502052951L, 457121784278309633L)), token);
        Assertions.assertNotNull(usersReply);
    }
}
