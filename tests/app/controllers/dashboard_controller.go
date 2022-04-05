package controllers

func (t *TestControllers) TestThatGetStatsWork() {

	_, err := testClient.GetRequest(ServerUrl + "status")
	t.AssertEqual(err, nil)
}
