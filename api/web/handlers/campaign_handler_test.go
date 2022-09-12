package handlers_test

// import (
// 	"net/web"
// 	"subscribers/domain/campaigns"
// 	"subscribers/helpers"
// 	"subscribers/helpers/fake"
// 	"subscribers/web/handlers"
// 	"testing"

// 	"github.com/brianvoe/gofakeit/v6"
// 	"github.com/stretchr/testify/assert"
// )

// func createNewCampaign(userId, userName string) campaigns.Campaign {
// 	entity := campaigns.NewCampaign(
// 		gofakeit.Sentence(3), gofakeit.Email(), gofakeit.Sentence(3), gofakeit.Sentence(6), userId, userName)
// 	fake.DB.Create(&entity)
// 	return *entity
// }

// func Test_campaign_post_validate_token(t *testing.T) {
// 	fake.Build()

// 	w := fake.MakeTestHTTP("POST", "/campaigns", nil, "")

// 	assert.Equal(t, web.StatusUnauthorized, w.Code)
// }

// func Test_campaign_post_validate_fields(t *testing.T) {
// 	fake.Build()

// 	w := fake.MakeTestHTTP("POST", "/campaigns", nil, fake.GenerateAnyToken())

// 	response := helpers.BufferToString(w.Body)
// 	assert.Contains(t, response, "'Name' is required")
// 	assert.Contains(t, response, "'From' is required")
// 	assert.Contains(t, response, "'Subject' is required")
// 	assert.Contains(t, response, "'Body' is required")
// }

// func Test_campaign_post_save_new_campaign(t *testing.T) {
// 	fake.Build()
// 	body := handlers.CampaignRequest{
// 		Name:    "teste 1",
// 		From:    "teste@teste.com.br",
// 		Subject: "Test 2",
// 		Body:    "Teste",
// 	}

// 	w := fake.MakeTestHTTP("POST", "/campaigns", body, fake.GenerateAnyToken())

// 	assert.Equal(t, web.StatusCreated, w.Code)
// }

// func Test_campaign_post_show_erro_when_not_create(t *testing.T) {
// 	fake.Build()
// 	mock := &fake.RepositoryMock[campaigns.Campaign]{
// 		ReturnsCreate: false,
// 	}
// 	fake.DI.CampaignHandler.CampaignRepository = mock
// 	body := handlers.CampaignRequest{
// 		Name:    "teste 1",
// 		From:    "teste@teste.com.br",
// 		Subject: "Test 2",
// 		Body:    "Teste",
// 	}

// 	w := fake.MakeTestHTTP("POST", "/campaigns", body, fake.GenerateAnyToken())

// 	assert.Equal(t, web.StatusInternalServerError, w.Code)
// }

// func Test_campaign_get_campaign_by_id(t *testing.T) {
// 	fake.Build()
// 	entity := createNewCampaign("xpto", gofakeit.Name())
// 	fake.DB.Create(&entity)

// 	w := fake.MakeTestHTTP("GET", "/campaigns/"+entity.ID, entity, fake.GenerateTokenWithUserId("xpto"))

// 	response := helpers.BufferToString(w.Body)
// 	assert.Contains(t, response, entity.Name)
// 	assert.Contains(t, response, entity.From)
// 	assert.Contains(t, response, entity.Body)
// 	assert.Equal(t, web.StatusOK, w.Code)
// }

// func Test_campaign_get_by_id_not_found(t *testing.T) {
// 	fake.Build()

// 	w := fake.MakeTestHTTP("GET", "/campaigns/id_invalid", nil, fake.GenerateAnyToken())

// 	response := helpers.BufferToString(w.Body)
// 	assert.Contains(t, response, "Not found")
// 	assert.Equal(t, web.StatusNotFound, w.Result().StatusCode)
// }

// func Test_campaign_get_campaign_by_id_must_calculate_dashboard(t *testing.T) {
// 	fake.Build()
// 	entity := createNewCampaign("xpto", gofakeit.Name())
// 	fake.DB.Create(&entity)
// 	sub1 := campaigns.NewSubscriber(entity, gofakeit.HexColor(), gofakeit.Email())
// 	fake.DB.Create(sub1)
// 	sub2 := campaigns.NewSubscriber(entity, gofakeit.HexColor(), gofakeit.Email())
// 	sub2.Sent()
// 	fake.DB.Create(sub2)
// 	sub3 := campaigns.NewSubscriber(entity, gofakeit.HexColor(), gofakeit.Email())
// 	sub3.Read()
// 	fake.DB.Create(sub3)

// 	w := fake.MakeTestHTTP("GET", "/campaigns/"+entity.ID, entity, fake.GenerateTokenWithUserId("xpto"))

// 	response := helpers.BufferToObj[handlers.CampaignResponse](w.Body)
// 	assert.Equal(t, 3, response.BaseOfSubscribers)
// 	assert.Equal(t, 1, response.TotalSent)
// 	assert.Equal(t, 1, response.TotalRead)
// }

// func Test_campaign_get_all_campaign_of_user(t *testing.T) {
// 	fake.Build()
// 	createNewCampaign("xpto", gofakeit.Name())
// 	createNewCampaign("xpto", gofakeit.Name())
// 	createNewCampaign("another_user_current", gofakeit.Name())
// 	amountOfCampaignsExpectedOfUser := 2

// 	w := fake.MakeTestHTTP("GET", "/campaigns", nil, fake.GenerateTokenWithUserId("xpto"))

// 	campaignsOfUser := helpers.BufferToObj[[]campaigns.Campaign](w.Body)
// 	assert.Equal(t, amountOfCampaignsExpectedOfUser, len(campaignsOfUser))
// }

// func Test_campaign_get_all_not_found(t *testing.T) {
// 	fake.Build()

// 	w := fake.MakeTestHTTP("GET", "/campaigns", nil, fake.GenerateAnyToken())

// 	response := helpers.BufferToString(w.Body)
// 	assert.Contains(t, response, "Not found")
// 	assert.Equal(t, web.StatusNotFound, w.Result().StatusCode)
// }
