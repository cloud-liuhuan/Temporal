package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RTradeLtd/Temporal/queue"
	"github.com/RTradeLtd/Temporal/rtfs"

	"github.com/RTradeLtd/Temporal/models"
	"github.com/jinzhu/gorm"

	"github.com/RTradeLtd/Temporal/utils"
	"github.com/gin-gonic/gin"
)

func PinToHostedIPFSNetwork(c *gin.Context) {
	cC := c.Copy()
	networkName, exists := cC.GetPostForm("network_name")
	if !exists {
		FailNoExist(c, "network_name post form does not exist")
		return
	}
	hash := cC.Param("hash")
	ethAddress := GetAuthenticatedUserFromContext(cC)
	holdTimeInMonths, exists := cC.GetPostForm("hold_time")
	if !exists {
		FailNoExist(c, "hold_time post form param not present")
		return
	}
	_, err := strconv.ParseInt(holdTimeInMonths, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, ok := cC.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}

	canUpload, err := CheckAccessForPrivateNetwork(ethAddress, networkName, db)
	if err != nil {
		FailOnError(c, err)
		return
	}
	if !canUpload {
		FailNotAuthorized(c, "unauthorized access to private network")
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	pnet, err := im.GetNetworkByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}

	apiURL := pnet.APIURL
	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		FailOnError(c, err)
		return
	}
	go func() {
		err := manager.Pin(hash)
		if err != nil {
			//TODO log and handle
			fmt.Println("Error encountered pinning content to private ipfs node", err.Error())
			return
		}
	}()
	c.JSON(http.StatusOK, gin.H{
		"status": "content pin request sent",
	})
}

//TODO: NEED TO FINISH
func GetFileSizeInBytesForObjectForHostedIPFSNetwork(c *gin.Context) {
	ethAddress := GetAuthenticatedUserFromContext(c)
	networkName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExistPostForm(c, "network_name")
		return
	}
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}
	canUpload, err := CheckAccessForPrivateNetwork(ethAddress, networkName, db)
	if err != nil {
		FailOnError(c, err)
		return
	}
	if !canUpload {
		FailNotAuthorized(c, "unauthorized access to private network")
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	key := c.Param("key")
	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	sizeInBytes, err := manager.GetObjectFileSizeInBytes(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"object":        key,
		"size_in_bytes": sizeInBytes,
	})

}

func AddFileToHostedIPFSNetwork(c *gin.Context) {
	ethAddress := GetAuthenticatedUserFromContext(c)
	if ethAddress != AdminAddress {
		FailNotAuthorized(c, "unauthorized access to private network upload")
		return
	}

	networkName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExist(c, "network_name post form does not exist")
		return
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}

	canUpload, err := CheckAccessForPrivateNetwork(ethAddress, networkName, db)
	if err != nil {
		FailOnError(c, err)
		return
	}
	if !canUpload {
		FailNotAuthorized(c, "unauthorized access to private network")
		return
	}
	holdTimeinMonths, exists := c.GetPostForm("hold_time")
	if !exists {
		FailNoExist(c, "hold_time post form does not exist")
		return
	}
	_, err = strconv.ParseInt(holdTimeinMonths, 10, 64)
	if err != nil {
		FailOnError(c, err)
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}

	ipfsManager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		FailOnError(c, err)
		return
	}

	fmt.Println("fetching file")
	// fetch the file, and create a handler to interact with it
	fileHandler, err := c.FormFile("file")
	if err != nil {
		FailOnError(c, err)
		return
	}

	file, err := fileHandler.Open()
	if err != nil {
		FailOnError(c, err)
		return
	}
	resp, err := ipfsManager.Shell.Add(file)
	if err != nil {
		FailOnError(c, err)
		return
	}
	fmt.Println("file uploaded")
	c.JSON(http.StatusOK, gin.H{
		"status": resp,
	})
}

//TODO: NEED TO FINISH
func IpfsPubSubPublishToHostedIPFSNetwork(c *gin.Context) {
	ethAddress := GetAuthenticatedUserFromContext(c)
	if ethAddress != AdminAddress {
		FailNotAuthorized(c, "unauthorized access to admin route")
		return
	}
	networkName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExistPostForm(c, "network_name")
		return
	}
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	topic := c.Param("topic")
	message, present := c.GetPostForm("message")
	if !present {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "message post form param is not present",
		})
		return
	}
	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = manager.PublishPubSubMessage(topic, message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"topic": topic, "message": message})
}

///TODO: NEED TO FINISH
func IpfsPubSubConsumeForHostedIPFSNetwork(c *gin.Context) {
	cC := c.Copy()

	ethAddress := GetAuthenticatedUserFromContext(cC)
	if ethAddress != AdminAddress {
		FailNotAuthorized(c, "unauthorized access to admin route")
		return
	}
	networkName, exists := cC.GetPostForm("network_name")
	if !exists {
		FailNoExistPostForm(c, "network_name")
		return
	}

	db, ok := cC.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to load database",
		})
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	topic := cC.Param("topic")

	go func() {
		manager, err := rtfs.Initialize("", apiURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		manager.SubscribeToPubSubTopic(topic)
		manager.ConsumeSubscription(manager.PubSub)
	}()

	c.JSON(http.StatusOK, gin.H{"status": "consuming messages in background"})
}

//TODO: NEED TO FINISH
func RemovePinFromLocalHostForHostedIPFSNetwork(c *gin.Context) {
	cC := c.Copy()
	ethAddress := GetAuthenticatedUserFromContext(cC)
	if ethAddress != AdminAddress {
		FailNotAuthorized(c, "unauthorized access to admin route")
		return
	}
	networkName, exists := cC.GetPostForm("network_name")
	if !exists {
		FailNoExistPostForm(c, "network_name")
		return
	}
	db, ok := cC.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	// fetch hash param
	hash := cC.Param("hash")

	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		FailOnError(c, err)
		return
	}
	// remove the file from the local ipfs state
	// TODO: implement some kind of error handling and notification
	err = manager.Shell.Unpin(hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error unpinning hash %s", err.Error()),
		})
		return
	}

	// TODO:
	// change to send a message to the cluster to depin
	mqConnectionURL := cC.MustGet("mq_conn_url").(string)
	qm, err := queue.Initialize(queue.IpfsQueue, mqConnectionURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO:
	// add in appropriate rabbitmq processing to delete from database
	qm.PublishMessage(hash)
	c.JSON(http.StatusOK, gin.H{"deleted": hash})
}

//TODO: NEED TO FINISH
func GetLocalPinsForHostedIPFSNetwork(c *gin.Context) {
	ethAddress := GetAuthenticatedUserFromContext(c)
	if ethAddress != AdminAddress {
		FailNotAuthorized(c, "unauthorized access to admin route")
		return
	}
	networkName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExistPostForm(c, "network_name")
		return
	}
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to load database",
		})
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	// initialize a connection toe the local ipfs node
	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		FailOnError(c, err)
		return
	}
	// get all the known local pins
	// WARNING: THIS COULD BE A VERY LARGE LIST
	pinInfo, err := manager.Shell.Pins()
	if err != nil {
		FailOnError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"pins": pinInfo})
}

// TODO: NEED TO FINISH
func GetObjectStatForIpfsForHostedIPFSNetwork(c *gin.Context) {
	networkName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExistPostForm(c, "network_name")
		return
	}
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	key := c.Param("key")
	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		FailOnError(c, err)
		return
	}
	stats, err := manager.ObjectStat(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

//TODO: NEED TO FINISH
func CheckLocalNodeForPinForHostedIPFSNetwork(c *gin.Context) {
	networkName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExistPostForm(c, "network_name")
		return
	}
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}
	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	hash := c.Param("hash")
	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	present, err := manager.ParseLocalPinsForHash(hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"present": present})
}

func PublishDetailedIPNSToHostedIPFSNetwork(c *gin.Context) {

	networkName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExist(c, "network_name post form does not exist")
		return
	}

	ethAddress := GetAuthenticatedUserFromContext(c)

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to load database",
		})
		return
	}
	canUpload, err := CheckAccessForPrivateNetwork(ethAddress, networkName, db)
	if err != nil {
		FailOnError(c, err)
		return
	}
	if !canUpload {
		FailNotAuthorized(c, "unauthorized access to private network")
		return
	}
	um := models.NewUserManager(db)

	hash, present := c.GetPostForm("hash")
	if !present {
		FailNoExistPostForm(c, "hash")
		return
	}
	lifetimeStr, present := c.GetPostForm("life_time")
	if !present {
		FailNoExistPostForm(c, "lifetime")
		return
	}
	ttlStr, present := c.GetPostForm("ttl")
	if !present {
		FailNoExistPostForm(c, "ttl")
		return
	}
	key, present := c.GetPostForm("key")
	if !present {
		FailNoExistPostForm(c, "key")
		return
	}
	resolveString, present := c.GetPostForm("resolve")
	if !present {
		FailNoExistPostForm(c, "resolve")
		return
	}

	ownsKey, err := um.CheckIfKeyOwnedByUser(ethAddress, key)
	if err != nil {
		FailOnError(c, err)
		return
	}

	if !ownsKey {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "attempting to generate IPNS entry unowned key",
		})
		return
	}

	im := models.NewHostedIPFSNetworkManager(db)
	apiURL, err := im.GetAPIURLByName(networkName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	manager, err := rtfs.Initialize("", apiURL)
	if err != nil {
		FailOnError(c, err)
		return
	}
	fmt.Println("creating key store manager")
	err = manager.CreateKeystoreManager()
	if err != nil {
		FailOnError(c, err)
		return
	}
	resolve, err := strconv.ParseBool(resolveString)
	if err != nil {
		FailOnError(c, err)
		return
	}
	lifetime, err := time.ParseDuration(lifetimeStr)
	if err != nil {
		FailOnError(c, err)
		return
	}
	ttl, err := time.ParseDuration(ttlStr)
	if err != nil {
		FailOnError(c, err)
		return
	}
	prePubTime := time.Now()
	keyID, err := um.GetKeyIDByName(ethAddress, key)
	fmt.Println("using key id of ", keyID)
	if err != nil {
		FailOnError(c, err)
		return
	}
	fmt.Println("publishing to IPNS")
	resp, err := manager.PublishToIPNSDetails(hash, key, lifetime, ttl, resolve)
	if err != nil {
		FailOnError(c, err)
		return
	}
	postPubTime := time.Now()
	timeDifference := postPubTime.Sub(prePubTime)

	ipnsManager := models.NewIPNSManager(db)
	ipnsEntry, err := ipnsManager.UpdateIPNSEntry(resp.Name, resp.Value, key, lifetime, ttl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"name":                   resp.Name,
		"value":                  resp.Value,
		"time_to_create_minutes": timeDifference.Minutes(),
		"ipns_entry_model":       ipnsEntry,
	})
}

func CreateHostedIPFSNetworkEntryInDatabase(c *gin.Context) {
	// lock down as admin route for now
	cC := c.Copy()
	ethAddress := GetAuthenticatedUserFromContext(cC)
	if ethAddress != AdminAddress {
		FailNotAuthorized(c, "unauthorized access")
		return
	}

	networkName, exists := cC.GetPostForm("network_name")
	if !exists {
		FailNoExist(c, "network_name post form does not exist")
		return
	}

	apiURL, exists := cC.GetPostForm("api_url")
	if !exists {
		FailNoExist(c, "api_url post form does not exist")
		return
	}

	swarmKey, exists := cC.GetPostForm("swarm_key")
	if !exists {
		FailNoExist(c, "swarm_key post form does not exist")
		return
	}

	bPeers, exists := cC.GetPostFormArray("bootstrap_peers")
	if !exists {
		FailNoExist(c, "boostrap_peers post form array does not exist")
		return
	}
	nodeAddresses, exists := cC.GetPostFormArray("local_node_addresses")
	if !exists {
		FailNoExist(c, "local_node_addresses post form array does not exist")
		return
	}
	users := cC.PostFormArray("users")
	var localNodeAddresses []string
	var bootstrapPeerAddresses []string

	if len(nodeAddresses) != len(bPeers) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "length of local_node_addresses and bootstrap_peers must be equal",
		})
		return
	}
	for k, v := range bPeers {
		addr, err := utils.GenerateMultiAddrFromString(v)
		if err != nil {
			FailOnError(c, err)
			return
		}
		valid, err := utils.ParseMultiAddrForIPFSPeer(addr)
		if err != nil {
			FailOnError(c, err)
			return
		}
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("provided peer %s is not a valid bootstrap peer", addr),
			})
			return
		}
		addr, err = utils.GenerateMultiAddrFromString(nodeAddresses[k])
		if err != nil {
			FailOnError(c, err)
			return
		}
		valid, err = utils.ParseMultiAddrForIPFSPeer(addr)
		if err != nil {
			FailOnError(c, err)
			return
		}
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("provided peer %s is not a valid ipfs peer", addr),
			})
			return
		}
		bootstrapPeerAddresses = append(bootstrapPeerAddresses, v)
		localNodeAddresses = append(localNodeAddresses, nodeAddresses[k])
	}
	// previously we were initializing like `var args map[string]*[]string` which was causing some issues.
	args := make(map[string][]string)
	args["local_node_peer_addresses"] = localNodeAddresses
	if len(bootstrapPeerAddresses) > 0 {
		args["bootstrap_peer_addresses"] = bootstrapPeerAddresses
	}
	db, ok := cC.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}
	manager := models.NewHostedIPFSNetworkManager(db)
	network, err := manager.CreateHostedPrivateNetwork(networkName, apiURL, swarmKey, args, users)
	if err != nil {
		FailOnError(c, err)
		return
	}
	um := models.NewUserManager(db)

	if len(users) > 0 {
		for _, v := range users {
			err := um.AddIPFSNetworkForUser(v, networkName)
			if err != nil {
				FailOnError(c, err)
				return
			}
		}
	} else {
		err := um.AddIPFSNetworkForUser(AdminAddress, networkName)
		if err != nil {
			FailOnError(c, err)
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{
		"network": network,
	})

}

func GetIPFSPrivateNetworkByName(c *gin.Context) {
	ethAddress := GetAuthenticatedUserFromContext(c)
	if ethAddress != AdminAddress {
		FailNotAuthorized(c, "unauthorized access")
		return
	}
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load database",
		})
		return
	}

	netName, exists := c.GetPostForm("network_name")
	if !exists {
		FailNoExist(c, "network_name post form does not exist")
		return
	}
	manager := models.NewHostedIPFSNetworkManager(db)
	net, err := manager.GetNetworkByName(netName)
	if err != nil {
		FailOnError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"network": net,
	})
}
