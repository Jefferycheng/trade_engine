# Trade Engine API

This document outlines the API specifications for the Trade Engine, a system that facilitates the matching and execution of buy and sell orders in a financial market.
The engine operates based on a web api, handling orders submitted by multiple traders concurrently.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Match Rules](#match-rules)
- [API Endpoints](#api-endpoints)
    - [Submit Order](#submit-order)
    - [Concurrent test](#concurrent-test)


## Introduction

The Trade Engine is designed to process buy and sell orders submitted by traders, matching orders with the same price based on their timestamp (First In, First Out - FIFO). Orders are executed in the order they are received.
The system supports market and limit orders, and multiple traders can execute orders concurrently.

## Getting Started

To interact with the Trade Engine, traders can use the provided API. The following sections outline the available endpoints and how to use them.


## Match Rules

#### Buy and limit order
Match the same price sell order from sell order queue

#### Buy and price order
Match the minimum price in the sell order queue

#### Sell and limit order
Match the same price buy order from buy order queue

#### Sell and price order
Match the maximum price in the buy order queue

## API Endpoints

### Submit Order

**Endpoint:** `GET /submit_order`

**Parameters:**
- `b` (Boolean): Buy (true) or sell (false).
- `p` (Integer): Price of the order. 0 indicates market price.
- `c` (Integer): Quantity of the asset to buy or sell.

**Example Request:**
```http
GET /submit_order?b=true&p=100&c=10
```

### Concurrent test

**Endpoint:** `GET /concurrent_test`

**Example Request:**
```http
GET /concurrent_test
```

