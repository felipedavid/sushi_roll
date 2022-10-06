
# Modelo de Dados

### Modelo Entidade-Relacionamento (MER) - Mermaid

```mermaid
    erDiagram 
            USER ||..o{ PROFESSIONAL_CRITIC : has

            USER ||..o{ USER_RATING : has

            USER ||..o{ COMMENT : has

            USER ||..o{ USER_ROLE : has

            GAME ||..o{ COMMENT : has

            GAME ||..o{ REFERRAL_LINK : has

            GAME ||..o{ USER_RATING : has

            GAME ||..o{ GAME_CATEGORY : has

            VENDOR ||..o{ REFERRAL_LINK : has

            ROLE ||..o{ USER_ROLE : has
            
            CATEGORY ||..o{ GAME_CATEGORY : has



            USER { 

            BIGSERIAL id 

            VARCHAR name 

            VARCHAR email 

            TIMESTAMPZ birth 

            TIMESTAMPZ createdAt 

            } 

            COMMENT { 

                BIGSERIAL id 

                VARCHAR content 

                BIGINT writerId 

                BIGINT gameId 

                TIMESTAMPZ createdAt 

            } 

            GAME{ 

                BIGSERIAL id 

                VARCHAR title 

                BIGSERIAL description 

                TIMESTAMPZ release 

                FLOAT avgUserRating 

                TIMESTAMPZ createdAt 

            } 

            USER_RATING{ 

                BIGSERIAL id 

                INT value 

                BIGINT userId 

                BIGINT gameId 

                TIMESTAMPZ createdAt 

            } 

            PROFESSIONAL_CRITIC { 

                BIGSERIAL id 

                BIGINT userId 

                FLOAT rating 

                VARCHAR commentary 

            }

            ROLE{
                BIGSERIAL id
                VARCHAR name
            }

            USER_ROLE{
                BIGINT userId
                BIGINT roleId
                TIMESTAMPZ createdAt
            }

            VENDOR{
                BIGSERIAL id
                VARCHAR name
                VARCHAR websiteUrl
                VARCHAR createdAt
            }

            REFERRAL_LINK{
                BIGSERIAL id
                BIGINT vendorId
                BIGINT gameId
                VARCHAR link
            }

            CATEGORY{
                BIGSERIAL id
                VARCHAR name
                INT ageGroup
            }

            GAME_CATEGORY{
                BIGINT gameId
                BIGINT categoryId
            } 	 
```
