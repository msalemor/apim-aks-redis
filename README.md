# Accessing services in AKS via API Management with Redis Cache


## Azure Services

- VNET
- AKS
- API Management
- Redis Cache
- Deploy service and deployment to AKS

## Deploy VNET

- Deploy a VNET: 10.0.0.0/16
- Deploy a AKS subnet 10.0.1.0/24 to the VNET
- Reserve subnet 10.0.2.0/24

> Note: AKS CNI requires two subnets

## Deploy AKS to VNET Subnet

```bash
az aks create \
    --resource-group <private-cluster-resource-group> \
    --name <private-cluster-name> \
    --load-balancer-sku standard \
    --enable-private-cluster \
    --network-plugin azure \
    --vnet-subnet-id <subnet-id> \
    --docker-bridge-address 172.17.0.1/16 \
    --dns-service-ip 10.2.0.10 \
    --service-cidr 10.2.0.0/24
```

## Deploy API Management (this deployment can take 45 minutes)

- From the Azure portal create an API Management

### Change API Management in external mode (this change can take 45 minutes)

> Note: in external mode API Management the portal and gateway are exposed to the internet but API Management can access internal services. In this example, the services will be hosted in AKS.

- Once API Management is created, modified network settings and change it to external
- Apply the settings

## Deploy Redis Cache

- From the portal, apply the settings

## Configure Redis Cache in API Management

- Get the primary connection string
- Configure it in Cache in API Management

## Configure the API in API Management

- GET to: http://x.x.x.x/api/contacts from: http://gateway.apim.com/contacts
- GET to: http://x.x.x.x/api/contacts/{state} with caching from http://gateway.apim.com/contacts

## Deploy service and deployment to AKS

```kubectl apply -f .```

## Testing
