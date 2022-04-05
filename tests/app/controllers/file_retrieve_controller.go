package controllers

func (t *TestControllers) TestThatRetrieveFileWorks() {

	file, err := testClient.GetRequest(ServerUrl + "files/test/test.file")
	t.AssertEqual(err, nil)
	t.AssertEqual(file, "this is a test file")
}
