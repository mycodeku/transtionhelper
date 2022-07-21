package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/mycodeku/transtionhelper/crypto/keys/secp256k1"
	"github.com/mycodeku/transtionhelper/simapp"
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/x/authz/keeper"
	bank "github.com/mycodeku/transtionhelper/x/bank/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper keeper.Keeper
}

func (suite *GenesisTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(suite.T(), checkTx)

	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1})
	suite.keeper = app.AuthzKeeper
}

var (
	granteePub  = secp256k1.GenPrivKey().PubKey()
	granterPub  = secp256k1.GenPrivKey().PubKey()
	granteeAddr = sdk.AccAddress(granteePub.Address())
	granterAddr = sdk.AccAddress(granterPub.Address())
)

func (suite *GenesisTestSuite) TestImportExportGenesis() {
	coins := sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(1_000)))

	now := suite.ctx.BlockTime()
	expires := now.Add(time.Hour)
	grant := &bank.SendAuthorization{SpendLimit: coins}
	err := suite.keeper.SaveGrant(suite.ctx, granteeAddr, granterAddr, grant, &expires)
	suite.Require().NoError(err)
	genesis := suite.keeper.ExportGenesis(suite.ctx)

	// TODO, recheck!
	// Clear keeper
	suite.keeper.DeleteGrant(suite.ctx, granteeAddr, granterAddr, grant.MsgTypeURL())

	suite.keeper.InitGenesis(suite.ctx, genesis)
	newGenesis := suite.keeper.ExportGenesis(suite.ctx)
	suite.Require().Equal(genesis, newGenesis)
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}