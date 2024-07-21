This micro-service is part of course-registration system.

This service manages profile and their corresponding data. <br>
Every successful  request would return <i>Status: 200</i> with requested data.<br>
In case of any error encoutered, <i>{response: "error message"} </i> would be returned as request's response with respective status code.

<h1>API Endpoints</h1>

<h3>Login</h3>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>POST</td>
    <td>/login</td>
    <td>Login using <i>email_id</i> and <i> password </i></td>
  </tr>
</table>

<h3>Admin</h3>
<hr>
<h4>Student actions:</h4>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>GET</td>
    <td>
      /students
      <br>
      /students?email_id=?
    </td>
    <td>Fetch all students or single profile based on their email_id</td>
  </tr>
  
  <tr>
    <td>POST</td>
    <td>
      /students
    </td>
    <td>Create a new student profile <br>
      <i>Email_id: primary key<br>First_name<br>Last_name<br>Program_enrolled</i>
    </td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/students/:email_id</td>
    <td>Update existing student profile data. <br>Fields to update: <br><i>First_name<br>Last_name<br>Program_enrolled</i></td>
  </tr>
  <tr>
    <td>DELTE</td>
    <td>/students/:email_id</td>
    <td>Delete existing student profile data based on email_id</td>
  </tr>
</table>

<h4>Professor actions:</h4>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>GET</td>
    <td>
      /professors
      <br>
      /professors?email_id=?
    </td>
    <td>Fetch all professors or single profile based on their email_id</td>
  </tr>
  
  <tr>
    <td>POST</td>
    <td>
      /professors
    </td>
    <td>Create a new professor profile <br>
      <i>Email_id: primary key<br>First_name<br>Last_name<br>Designation<br>Department</i>
    </td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/professors/:email_id</td>
    <td>Update existing professor profile data. <br>Fields to update: <br><i>First_name<br>Last_name<br>Department<br>Designation</i></td>
  </tr>
  <tr>
    <td>DELTE</td>
    <td>/professors/:email_id</td>
    <td>Delete existing professor profile data based on email_id</td>
  </tr>
</table>

<h3>Student</h3>
<hr>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  
  <tr>
    <td>GET</td>
    <td>/:email_id</td>
    <td>Get student profile <br> Field<br><i>email_id</i></td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/password/:email_id</td>
    <td>Update password for a student profile<br> Field<br><i>new_password</i></td>
  </tr>
</table>

<h3>Professor</h3>
<hr>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  
  <tr>
    <td>GET</td>
    <td>/:email_id</td>
    <td>Get professor profile <br> Field<br><i>email_id</i></td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/password/:email_id</td>
    <td>Update password for a professor profile<br> Field<br><i>new_password</i></td>
  </tr>
</table>
