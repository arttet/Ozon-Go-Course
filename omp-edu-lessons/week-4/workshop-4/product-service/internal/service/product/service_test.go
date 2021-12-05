package product_service

//
//func setup(t *testing.T) (*Service, *MockIRepository, *MockICategoryClient) {
//	ctrl := gomock.NewController(t)
//	repo := NewMockIRepository(ctrl)
//	client := NewMockICategoryClient(ctrl)
//
//	service := &Service{
//		repo:   repo,
//		client: client,
//	}
//
//	return service, repo, client
//}
//
//func TestCreateProduct_Success_ReturnName(t *testing.T) {
//	service, repo, client := setup(t)
//
//	repo.EXPECT().
//		SaveProduct(gomock.Any(), gomock.Any()).
//		DoAndReturn(func(ctx context.Context, product *Product) error {
//			product.ID = 124
//			return nil
//		})
//
//	client.EXPECT().
//		IsCategoryExists(gomock.Any(), int64(4)).
//		Return(true, nil)
//
//	product, err := service.CreateProduct(
//		context.Background(),
//		"new product",
//		4,
//	)
//
//	require.Nil(t, err)
//	require.Equal(t, "new product", product.Name)
//}
//
//func TestCreateProduct_Success_ReturnCategoryID(t *testing.T) {
//	service, repo, client := setup(t)
//
//	repo.EXPECT().
//		SaveProduct(gomock.Any(), gomock.Any()).
//		DoAndReturn(func(ctx context.Context, product *Product) error {
//			product.ID = 124
//			return nil
//		})
//
//	client.EXPECT().
//		IsCategoryExists(gomock.Any(), int64(25)).
//		Return(true, nil)
//
//	product, err := service.CreateProduct(
//		context.Background(),
//		"new product",
//		25,
//	)
//
//	require.Nil(t, err)
//	require.Equal(t, int64(25), product.CategoryID)
//}
//
//func TestCreateProduct_Success_ReturnID(t *testing.T) {
//	service, repo, client := setup(t)
//
//	repo.EXPECT().
//		SaveProduct(gomock.Any(), gomock.Any()).
//		DoAndReturn(func(ctx context.Context, product *Product) error {
//			product.ID = 124
//			return nil
//		})
//
//	client.EXPECT().
//		IsCategoryExists(gomock.Any(), int64(3)).
//		Return(true, nil)
//
//	product, err := service.CreateProduct(
//		context.Background(),
//		"another product",
//		3,
//	)
//
//	require.Nil(t, err)
//	require.Equal(t, int64(124), product.ID)
//}
//
//func TestCreateProduct_CategoryDoesNotExist(t *testing.T) {
//	service, _, client := setup(t)
//
//	client.EXPECT().
//		IsCategoryExists(gomock.Any(), int64(3)).
//		Return(false, nil)
//
//	_, err := service.CreateProduct(
//		context.Background(),
//		"another product",
//		3,
//	)
//
//	require.NotNil(t, err)
//}
